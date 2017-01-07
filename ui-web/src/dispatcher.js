'use strict';

import Riot from 'riot';

let dispatcher = {
    // Switch page events
    PAGE_CHANGE: 'PAGE_CHANGE',
    PAGE_CHANGED: 'PAGE_CHANGED',

    TRY_LOGIN: 'TRY_LOGIN',
    RE_LOGIN: 'RE_LOGIN',
    SHOW_LOGIN_DIALOG: 'SHOW_LOGIN_DIALOG',
    DO_LOGOUT: 'DO_LOGOUT',
    PROFILE_CHANGED: 'PROFILE_CHANGED',
    SEARCH_APPOINTMENTS: 'SEARCH_APPOINTMENTS',
    SEARCH_APPOINTMENTS_CHANGED: 'SEARCH_APPOINTMENTS_CHANGED',
    NEW_APPOINTMENT: 'NEW_APPOINTMENT',
    GET_APPOINTMENT: 'GET_APPOINTMENT',
    SAVE_APPOINTMENT: 'SAVE_APPOINTMENT',
};

Riot.observable(dispatcher);

module.exports = dispatcher;
