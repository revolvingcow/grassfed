(function () {
    var po = document.createElement('script');
    po.type = 'text/javascript';
    po.async = true;
    po.src = 'https://apis.google.com/js/client:plusone.js';

    var s = document.getElementsByTagName('script')[0];
    s.parentNode.insertBefore(po, s);
})();

function onAuthentication (result) {
    if (result['status']['signed_in']) {
        $('div#authentication').stop().hide();
        gapi.client.load('plus', 'v1', apiClientLoaded);
    }
    else {
        if (console) {
            console.log('Sign-in state: ' + result['error']);
        }
    } 
}

function apiClientLoaded () {
    gapi.client.plus.people
        .get({ userId: 'me' })
        .execute(function (response) {
            // Determine if this account is already registered.
            $.post('/me/logon', { id: response.id })
            .done(function (response) {
                if (response) {
                    $("div#profile").show();
                }
            })
            .fail(function () {
            });
        });
}
