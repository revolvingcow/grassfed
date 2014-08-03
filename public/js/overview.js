$(function () {
    $.post('/overview')
    .done(function (response) {
        if (response) {
            $('.accounts').text(response.Accounts);
            $('.calories').text(response.Calories);
        }
    });
});
