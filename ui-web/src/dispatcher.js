'use strict';

import Riot from 'riot';

let dispatcher = {
    // Switch page events
    PAGE_CHANGE: 'PAGE_CHANGE',
    PAGE_CHANGED: 'PAGE_CHANGED',

    TRY_LOGIN: 'TRY_LOGIN',
    SHOW_LOGIN_DIALOG: 'SHOW_LOGIN_DIALOG',
    DO_LOGOUT: 'DO_LOGOUT',
    PROFILE_CHANGED: 'PROFILE_CHANGED',
    SEARCH_PATIENTS: 'SEARCH_PATIENTS',
    SEARCH_PATIENTS_CHANGED: 'SEARCH_PATIENTS_CHANGED',
    GET_PATIENT_DETAILS: 'GET_PATIENT_DETAILS',
};

Riot.observable(dispatcher);

module.exports = dispatcher;
