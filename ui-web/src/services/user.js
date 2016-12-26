'use strict';

import dispatcher from '../dispatcher';
import _fetch from '../util/fetch';
import respHelper from '../util/response-helper';
import cookies from 'cookie';

dispatcher.on(dispatcher.TRY_LOGIN, _login);
dispatcher.on(dispatcher.DO_LOGOUT, _logout);

function _login() {
    dispatcher.trigger(dispatcher.SHOW_LOGIN_DIALOG, {});
}

const authCookieName = 'X-App-Auth';

function _logout() {
    document.cookie = cookies.serialize(authCookieName, "", {expires: new Date(0)});
    dispatcher.trigger(
        dispatcher.PROFILE_CHANGED, {
            login: '',
            username: '',
            authorized: false
        });
}

(window.onpopstate = function () {
    _fetch.setAuthCookieName(authCookieName);
    let token = cookies.parse(document.cookie)[authCookieName];
    if (token) {
        _fetch.fetch('/api/profile')
            .then(r => {
                return respHelper.handleStatus(r);
            })
            .then(data => {
                const {error} = data;
                if (error) {
                    console.log(error);
                } else {
                    data.authorized = true;
                    dispatcher.trigger(dispatcher.PROFILE_CHANGED, data);
                }
            })
            .catch(e => {
                console.log(e);
                if (e.error === 'AUTH') {
                    alert('Auth error. Try to re-login');
                    _logout();
                }
            });
    }
})();
