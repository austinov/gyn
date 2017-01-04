'use strict';

import dispatcher from '../dispatcher';
import _fetch from '../utils/fetch';
import respHelper from '../utils/response-helper';

dispatcher.on(dispatcher.SEARCH_PATIENTS, (params) => {
	_searchPatients(params);
});

dispatcher.on(dispatcher.NEW_APPOINTMENT, _newAppointment);
dispatcher.on(dispatcher.GET_APPOINTMENT, _getAppointment);
dispatcher.on(dispatcher.SAVE_APPOINTMENT, _saveAppointment);

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

function _newAppointment(cb) {
    if (cb) {
        _fetch.fetch('/api/dictionaries')
            .then(r => {
                return respHelper.handleStatus(r);
            })
            .then(data => {
                const {error} = data;
                if (error) {
                    console.log(error);
                } else {
                   cb(data, _getDefaultData())
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
}

function _getAppointment(id, cb) {
    // TODO fetch
    if (cb) {
        Promise.all([
            _fetch.fetch('/api/dictionaries'),
            _fetch.fetch('/api/appointment/' + id)
        ])
        .then(responses => {
            return Promise.all([
                    respHelper.handleStatus(responses[0]),
                    respHelper.handleStatus(responses[1])]);
        })
        .then(data => {
            let obj1 = data[0],
                obj2 = data[1];
            console.log('====.1', obj1);
            console.log('====.2', obj2);
            const {error} = obj2;
            if (error) {
                 console.log('>>>> Error', error);
            }
        })
        .catch(e => {
            console.log('Errorroorr', e);
            cb({}, {}, e);
        });
        /*
        _fetch.fetch('/api/dictionaries')
            .then(r => {
                return respHelper.handleStatus(r);
            })
            .then(data => {
                const {error} = data;
                if (error) {
                    console.log(error);
                } else {
                   // TODO patient data
                   cb(data, {
                        id: id,
                        name: 'Иванова Мария Ивановна',
                        health_state_id: 3,
                        skin_state_id: 2
                   });
                }
            })
            .catch(e => {
                console.log(e);
                if (e.error === 'AUTH') {
                    _logout();
                    dispatcher.trigger(dispatcher.SHOW_LOGIN_DIALOG, {});
                }
            });
        */
    }
}

function _saveAppointment(data, cb) {
    // TODO
    if (cb) {
        console.log('save', data);
        cb('error msg');
    }
}

const userName = 'userName';

function _getDefaultData() {
    let birthPlan = '- роды вести per vias naturals под адекватным обезболиванием (ЭДА)\n';
    birthPlan += '- профилактика кровотечения в  III и раннем послеродовом периодах\n';
    birthPlan += '- функциональная оценка таза\n';
    birthPlan += '- в случае функционального ухудшения плода или отклонении от нормального течения родов своевременно решить вопрос об оперативном родоразрешении';
    return {
        dateReceipt: _getToday(),
        doctorName: localStorage.getItem(userName),
        howReceipt: 'самотеком',
        alergo: 'отр',
        contactInfectied: 'отр',
        hiv: 'отр',
        transfusion: 'отр',
        dyscountry: 'отр',
        smoking: 'отр',
        drugs: 'отр',
        inheritance: 'не отягощена',
        head: 'не болит',
        vision: 'ясное',
        pulseType: 'ритмичный',
        throat: 'чистый',
        peritoneal: 'отрицательные',
        labors: 'abs',
        arches: 'свободные',
        birthPlan: birthPlan
    }
}

function _getToday() {
  let today = new Date();
  let dd = today.getDate();
  let mm = today.getMonth()+1; //January is 0!
  let yyyy = today.getFullYear();
  if(dd<10) {
      dd = '0' + dd
  } 
  if(mm < 10) {
      mm = '0' + mm
  } 
  return yyyy + '-' + mm +'-' + dd;
}
