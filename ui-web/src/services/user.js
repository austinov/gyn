'use strict';

import dispatcher from '../dispatcher';
import _fetch from '../utils/fetch';
import respHelper from '../utils/response-helper';
import cookies from 'cookie';

dispatcher.on(dispatcher.TRY_LOGIN, _login);
dispatcher.on(dispatcher.RE_LOGIN, _relogin);
dispatcher.on(dispatcher.DO_LOGOUT, _logout);

function _login() {
    dispatcher.trigger(dispatcher.SHOW_LOGIN_DIALOG, {});
}

const authCookieName = 'X-App-Auth';
const userName = 'userName';

function _logout() {
    document.cookie = cookies.serialize(authCookieName, "", {expires: new Date(0)});
    localStorage.removeItem(userName);
    dispatcher.trigger(
        dispatcher.PROFILE_CHANGED, {
            login: '',
            username: '',
            authorized: false
        });
}

function _relogin() {
    _logout();
    dispatcher.trigger(dispatcher.SHOW_LOGIN_DIALOG, {});
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
                    localStorage.setItem(userName, data.username);
                    dispatcher.trigger(dispatcher.PROFILE_CHANGED, data);
                }
            })
            .catch(e => {
                console.log(e);
                if (e.error === 'AUTH') {
                    _logout();
                    dispatcher.trigger(dispatcher.SHOW_LOGIN_DIALOG, {});
                }
            });
    }
})();
