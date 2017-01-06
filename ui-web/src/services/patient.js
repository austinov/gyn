'use strict';

import dispatcher from '../dispatcher';
import _fetch from '../utils/fetch';
import respHelper from '../utils/response-helper';
//import date from '../utils/date';

dispatcher.on(dispatcher.SEARCH_APPOINTMENTS, _searchAppointments);
dispatcher.on(dispatcher.NEW_APPOINTMENT, _newAppointment);
dispatcher.on(dispatcher.GET_APPOINTMENT, _getAppointment);
dispatcher.on(dispatcher.SAVE_APPOINTMENT, _saveAppointment);

function _searchAppointments(params, cb) {
    // TODO dateReceipt
	console.log('_searchAppointments', params);
    if (cb) {
        _fetch.fetch('/api/appointments', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(params)
            })
            .then(r => {
                return respHelper.handleStatus(r);
            })
            .then(data => {
                const {error} = data;
                if (error) {
                    console.log(error);
                }
                cb(data, error);
            })
            .catch(e => {
                console.log(e);
                if (e.error === 'AUTH') {
                    alert('Ошибка проверки пользователя.\nПопробуйте выйти/войти в приложение.');
                    _logout();
                    dispatcher.trigger(dispatcher.SHOW_LOGIN_DIALOG, {});
                } else {
                    cb([], e);
                }
            });
    }
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
        .then(r => {
            return Promise.all([
                    respHelper.handleStatus(r[0]),
                    respHelper.handleStatus(r[1])]);
        })
        .then(resp => {
            let dict = resp[0],
                data = resp[1];
            console.log('====.1', dict);
            console.log('====.2', data);
            const {error} = data;
            if (error) {
                 console.log('>>>> Error', error);
            }
            cb(dict, data, error);
        })
        .catch(e => {
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
    if (cb) {
        _fetch.fetch('/api/appointment', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            .then(r => {
                return respHelper.handleStatus(r);
            })
            .then(resp => {
                const {error} = resp;
                if (error) {
                    console.log(error);
                } else {
                    data.id = resp.id;
                }
                cb(data.id, error);
            })
            .catch(e => {
                console.log(e);
                if (e.error === 'AUTH') {
                    alert('Ошибка проверки пользователя при сохранении данных.\nПопробуйте выйти/войти в приложение.');
                    _logout();
                    dispatcher.trigger(dispatcher.SHOW_LOGIN_DIALOG, {});
                } else {
                    cb(null, e);
                }
            });
    }
}

const userName = 'userName';

function _getDefaultData() {
    let birthPlan = '- роды вести per vias naturals под адекватным обезболиванием (ЭДА)\n';
    birthPlan += '- профилактика кровотечения в  III и раннем послеродовом периодах\n';
    birthPlan += '- функциональная оценка таза\n';
    birthPlan += '- в случае функционального ухудшения плода или отклонении от нормального течения родов своевременно решить вопрос об оперативном родоразрешении';
    return {
        dateReceipt: new Date()/1000,
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
        dysuric: false,
        bowel: true,
        arches: 'свободные',
        birthPlan: birthPlan
    }
}
