'use strict';

import dispatcher from '../dispatcher';
import _fetch from '../util/fetch';
import respHelper from '../util/response-helper';

dispatcher.on(dispatcher.SEARCH_PATIENTS, (params) => {
	_searchPatients(params);
});

dispatcher.on(dispatcher.GET_PATIENT_DETAILS, _getPatientDetails);

function _searchPatients(params) {
    // TODO fetch
	console.log('_searchPatients', params);
    dispatcher.trigger(dispatcher.SEARCH_PATIENTS_CHANGED, [
        {
            id: 111,
            name: 'Иванова Мария Ивановна',
            date: '01.12.2016',
        },
        {
            id: 222,
            name: 'Петрова Ольга Васильевна',
            date: '02.12.2016',
        },
        {
            id: 333,
            name: 'Сидорова Светлана Николаевна',
            date: '03.12.2016',
        },
    ]);
}

function _getPatientDetails(id, cb) {
    // TODO fetch
    if (cb) {
        cb({
             id: id,
             name: 'Иванова Мария Ивановна'
        });
    }
}

