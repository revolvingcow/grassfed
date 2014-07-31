$(function () {
    var chart = $("#goalChart")[0];

    if (chart) {
        var ctx = chart.getContext("2d");
        var doughnutChart = new Chart(ctx).Doughnut(
            [
                {
                    value: getCurrentCalories(),
                    color: "#46bfbd",
                    highlight: "#46bfbd",
                    label: "Today"
                },
                {
                    value: getDailyCalories() - getCurrentCalories(),
                    color: "#ddd",
                    highlight: "#ddd",
                    label: "Remaining"
                }
            ],
            {
                animateScale: true
            });
    }

    function getDailyCalories() {
        return parseInt($('input[name="calories"][type="range"]').val());
    }

    function getCurrentCalories() {
        var calories = 0;
        return calories;
    }

    $(document).on('change', 'input[name="calories"][type="range"]', function (e) {
        var goal = getDailyCalories();
        var count = getCurrentCalories();
        var goalColor = "#ddd";
        var countColor = "#46bfbd";

        if (count >= goal) {
            countColor = "#f7464a";
        }

        $('span.goal').text(goal);

        if (doughnutChart) {
            console.log(doughnutChart.segments);

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
    });
});
