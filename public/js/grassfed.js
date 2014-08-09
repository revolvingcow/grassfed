(function ($) {
    $.each(['show', 'hide'], function (i, ev) {
        var el = $.fn[ev];
        $.fn[ev] = function () {
            this.trigger(ev);
            return el.apply(this, arguments);
        };
    });
})(jQuery);

$(function () {
    var chart = $('#goalChart')[0];
    var trends = $('#trendsChart')[0];
    var doughnutChart;
    var lineChart;
    var trendData;

    function startEngine() {
        // Pull the information we need first.
        loadStatistics();
        loadTrends();
        loadHistory();

        // Set focus.
        $('input[name="products"][type="text"]').focus();
    }

    // Check to see if the profile is not hidden. If it is not then start your engines!
    if (!$('div#profile').is(':hidden')) {
        startEngine();
    }
    else {
        $('div#profile').on('show', function () {
            startEngine();
        });
    }

    $('form#entry').on('submit', function (e) {
        e.preventDefault();

        // Get the current values.
        var product = $('input[name="product"][type="text"]').val();
        var calories = parseInt($('input[name="calories"][type="number"]').val());

        $.post('/me/add', { 'product': product, 'calories': calories })
            .done(function (response) {
                if (response) {
                    // Reset the values.
                    $('input[name="product"][type="text"]').val('');
                    $('input[name="calories"][type="number"]').val(0);
               
                    // Record the moment in history.
                    prependProduct(response.Id, response.Product, response.Calories);

                    // Update the chart.
                    var previousCalories = parseInt($('input[name="current"][type="hidden"]').val());
                    $('input[name="current"][type="hidden"]').val(previousCalories + response.Calories);
                    updateCalorieChart();
                    loadTrends();

                    // Reset focus.
                    $('input[name="products"][type="text"]').focus();
                }
            });

        return false;
    });

    $('form#recordWeight').on('submit', function (e) {
        e.preventDefault();

        var weight = $('input[name="weight"][type="number"]').val();

        $.post('/me/weight', { 'weight': weight })
            .done(function (response) {
                if (response) {
                    // Reset the values.
                    $('input[name="weight"][type="number"]').val('');
                    loadTrends();
                }
            });
    });

    $(document).on('click', 'button.delete-history', function (e) {
        e.preventDefault();
        var moment = parseInt($(this).parents('div.media').attr('data-moment'));

        $.ajax({
            url: '/me/history/' + moment,
            type: 'DELETE'
        })
        .done(function (response) {
            if (response) {
                var media = $('div.media[data-moment="' + moment + '"]');
                var panel = $(media).parents('div.panel');
                var previousCalories = parseInt($('input[name="current"][type="hidden"]').val());
                var calories = parseInt($(media).find('span.calories').text());
               
                // Remove the current media element which has been trashed.
                $(media).remove();

                // Update our calorie chart.
                $('input[name="current"][type="hidden"]').val(previousCalories - calories);
                updateCalorieChart();
                loadTrends();

                // Check if there are any more media elements in the panel. If not remove the panel.
                if ($(panel).find('div.media').length == 0) {
                    $(panel).remove();
                }
            }
        });
    });

    $(document).on('change', 'input[name="calories"][type="range"]', function (e) {
        var goal = $(this).val();
        
        // Send the new goal to the server.
        setGoal(goal);

        // Update the text for immediate feedback.
        $('span.goal').text(goal);

        // Update the chart.
        updateCalorieChart();
        loadTrends();
    });

    function prependProduct(id, product, calories) {
        var momentDate = new Date().toDateString();
        var panelTitle = $('h3.panel-title:contains("' + momentDate + '")');
        var panelBody;

        if (panelTitle.length > 0) {
            panelBody = $(panelTitle).parents('div.panel').find('div.panel-body');
        }
        else {
            $('div#history').prepend(
                $('<div class="panel panel-default">')
                    .append('<div class="panel-heading"><h3 class="panel-title">' + momentDate + '</h3></div>')
                    .append('<div class="panel-body">'));

            panelBody = $('h3.panel-title:contains("' + momentDate + '")').parents('div.panel').find('div.panel-body');
        }

        $(panelBody).prepend(
            $('<div class="media" data-moment="' + id + '">')
                .append('<div class="media-object pull-right"><button class="btn btn-danger delete-history"><span class="glyphicon glyphicon-fire"></span></button></div>')
                .append('<div class="media-body">')
                .append('<h4 class="media-heading">' + product + ' (<span class="calories">' + calories + '</span> calories)</h4>'));
    }

    function getDailyCalories() {
        return parseInt($('input[name="calories"][type="range"]').val());
    }

    function getCurrentCalories() {
        return parseInt($('input[name="current"][type="hidden"]').val());
    }
    
    function setGoal(calories) {
        $.post('/me/goal', { 'calories': calories });
    }

    function loadStatistics() {
        $.get('/me/stats')
            .done(function (response) {
                if (response) {
                    $('input[name="calories"][type="range"]').val(response.Goal);
                    $('span.goal').text(response.Goal);
                    $('input[name="current"][type="hidden"]').val(response.Current);
                    $('.streak').text(response.Streak);
                    $('.streak-unit').text('days');
                }

                updateCalorieChart();
            });
    }

    function loadTrends() {
        $.get('/me/trends')
            .done(function (response) {
                if (response) {
                    goals = [];
                    calories = [];
                    weights = [];

                    for (var i = 0; i < response.Labels.length; i++) {
                        goals.push(response.Goals[i]);
                        calories.push(response.History[i]);
                        weights.push(response.Weights[i]);
                    }

                    updateTrendChart(
                        response.Labels, 
                        goals, 
                        calories, 
                        weights);
                }
            });
    }

    function loadHistory() {
        $.get('/me/history')
            .done(function (response) {
                if (response) {
                    var history = $('div#history');
                    var moments = [];
                    var lastDate;
                    var panel;

                    for (var i = 0; i < response.length; i++) {
                        if (response[i]) {
                            var momentDate = new Date(response[i].Date).toDateString();
                            var moment = $('<div class="media" data-moment="' + response[i].Id + '">')
                                .append('<div class="media-object pull-right"><button class="btn btn-danger delete-history"><span class="glyphicon glyphicon-fire"></span></button></div>')
                                .append('<div class="media-body">')
                                .append('<h4 class="media-heading">' + response[i].Product + ' (<span class="calories">' + response[i].Calories + '</span> calories)</h4>');

                            if (!lastDate || lastDate != momentDate) {
                                lastDate = momentDate;
    
                                if (panel) {
                                    $(panel).find('div.panel-body').append(moments);
                                    $(history).append(panel);
                                    moments = [];
                                    panel = null;
                                }
    
                                panel = $('<div class="panel panel-default">')
                                    .append('<div class="panel-heading"><h3 class="panel-title">' + lastDate + '</h3></div>')
                                    .append('<div class="panel-body">');
                            }
                                
                            moments.push(moment);
                        }
                    }

                    if (panel) {
                        $(panel).find('div.panel-body').append(moments);
                        $(history).append(panel);
                    }
                }
            });
    }

    function updateTrendChart(labels, goals, calories, weights) {
        //if (!lineChart) {
            var ctxTrends = trends.getContext("2d");
            lineChart = new Chart(ctxTrends).Line(
                {
                    labels: labels, 
                    datasets: [
                        {
                            label: 'Goal',
                            fillColor: 'rgba(220, 220, 220, 0.2)',
                            strokeColor: 'rgba(220, 220, 220, 1)',
                            pointColor: 'rgba(220, 220, 220, 1)',
                            pointStrokeColor: '#fff',
                            pointHighlightFill: '#fff',
                            pointHighlightStroke: 'rgba(220, 220, 220, 1)',
                            data: goals 
                        },
                        {
                            label: 'Calories',
                            fillColor: 'rgba(196, 46, 42, 0.2)',
                            strokeColor: 'rgba(196, 46, 42, 0.5)',
                            pointColor: 'rgba(220, 220, 220, 1)',
                            pointStrokeColor: '#fff',
                            pointHighlightFill: '#fff',
                            pointHighlightStroke: 'rgba(220, 220, 220, 1)',
                            data: calories 
                        },
                        {
                            label: 'Weight',
                            fillColor: 'rgba(51, 51, 204, 0.2)',
                            strokeColor: 'rgba(51, 51, 204, 0.5)',
                            pointColor: 'rgba(51, 51, 204, 1)',
                            pointStrokeColor: '#fff',
                            pointHighlightFill: '#fff',
                            pointHighlightStroke: 'rgba(220, 220, 220, 1)',
                            data: weights 
                        }
                    ]
                },
                {
                    animateScale: true,
                    pointDot: false
                });
        //}
    };

    function updateCalorieChart() {
        var goal = getDailyCalories();
        var count = getCurrentCalories();
        var goalColor = "#ddd";
        var countColor = "#46bfbd";

        if (count >= goal) {
            countColor = "#f7464a";
        }
        
        if (!doughnutChart) {
            var ctxCalories = chart.getContext("2d");
            doughnutChart = new Chart(ctxCalories).Doughnut(
                [
                    {
                        value: count, 
                        color: "#46bfbd",
                        highlight: "#46bfbd",
                        label: "Today"
                    },
                    {
                        value: Math.max(goal - count, 0),
                        color: "#ddd",
                        highlight: "#ddd",
                        label: "Remaining"
                    }
                ],
                {
                    animateScale: true
                });
        }

        if (doughnutChart) {
            doughnutChart.segments[0].value = count;
            doughnutChart.segments[0].color = countColor;
            doughnutChart.segments[0].fillColor = countColor;
            doughnutChart.segments[0].highlight = countColor;
            
            doughnutChart.segments[1].value = Math.max(goal - count, 0);
            doughnutChart.segments[1].color = goalColor;
            doughnutChart.segments[1].fillColor = goalColor;
            doughnutChart.segments[1].highlight = goalColor;

            doughnutChart.update();
        }
    }
});
