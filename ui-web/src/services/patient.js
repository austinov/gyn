'use strict';

import dispatcher from '../dispatcher';
import _fetch from '../util/fetch';
import respHelper from '../util/response-helper';

dispatcher.on(dispatcher.SEARCH_PATIENTS, (params) => {
	_searchPatients(params);
});

function _searchPatients(params) {
	console.log('_searchPatients', params);
    dispatcher.trigger(dispatcher.SEARCH_PATIENTS_CHANGED, [
        {
            name: 'Иванова Мария Ивановна',
            date: '01.12.2016',
        },
        {
            name: 'Петрова Ольга Васильевна',
            date: '02.12.2016',
        },
        {
            name: 'Сидорова Светлана Николаевна',
            date: '03.12.2016',
        },
    ]);
}
