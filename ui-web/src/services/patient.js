'use strict';

import dispatcher from '../dispatcher';
import _fetch from '../utils/fetch';
import respHelper from '../utils/response-helper';

dispatcher.on(dispatcher.SEARCH_APPOINTMENTS, _searchAppointments);
dispatcher.on(dispatcher.NEW_APPOINTMENT, _newAppointment);
dispatcher.on(dispatcher.GET_APPOINTMENT, _getAppointment);
dispatcher.on(dispatcher.GET_APPOINTMENT_DOCX, _getAppointmentDocx);
dispatcher.on(dispatcher.SAVE_APPOINTMENT, _saveAppointment);

function _searchAppointments(params, cb) {
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
                _catchError(e, false, () => {
                    cb([], e);
                });
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
                _catchError(e, true);
            });
    }
}

function _getAppointment(id, cb) {
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
            const {error} = data;
            if (error) {
                 console.log(error);
            }
            cb(dict, data, error);
        })
        .catch(e => {
            _catchError(e, true, () => {
                cb({}, {}, e);
            });
        });
    }
}

function _getAppointmentDocx(id) {
    _fetch.fetch('/api/appointment/' + id + '/docx')
        .then(r => {
            return respHelper.handleStatus(r, true);
        })
        .then(data => {
            const {error} = data;
            if (error) {
                console.log(error);
            } else {
                let link = document.createElement('a');
                link.href = window.URL.createObjectURL(data.blob);
                link.download = data.filename;
                link.click();
            }
        })
        .catch(e => {
            _catchError(e, true);
        });
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
                _catchError(e, true, () => {
                    cb([], e);
                });
            });
    }
}

const userName = 'userName';

function _getDefaultData() {
    let conclusion = '1) госпитализируется в АОПБ согласно приказу №572н \n';
    conclusion += '2) динамическое наблюдение в условиях АОПБ\n';
    conclusion += '3) АД, УЗИ, допплерометрия – \n';
    conclusion += '4) Терапия, направленная на ';

    let birthPlan = '- роды вести per vias naturals под адекватным обезболиванием (ЭДА)\n';
    birthPlan += '- профилактика кровотечения в  III и раннем послеродовом периодах\n';
    birthPlan += '- функциональная оценка таза\n';
    birthPlan += '- в случае функционального ухудшения плода или отклонении от нормального ';
    birthPlan += 'течения родов своевременно решить вопрос об оперативном родоразрешении';
    return {
        dateReceipt: new Date()/1000,
        doctorName: localStorage.getItem(userName),
        alergo: 'отр',
        contactInfected: 'отр',
        hiv: 'отр',
        transfusion: 'отр',
        dyscountry: 'отр',
        smoking: 'отр',
        drugs: 'отр',
        history: 'ОРВИ',
        inheritance: 'не отягощена',
        paritet: 'Данная беременность наступила',
        firstTrimester: 'Скрининг - ',
        tromboflebiaStateId: 3,
        head: 'не болит',
        vision: 'ясное',
        lymph: 'не пальпируются',
        pulseType: 'ритмичный',
        throat: 'чистый',
        peritoneal: 'отрицательные',
        labors: 'abs',
        dysuric: false,
        bowel: true,
        cervixChannel: ' п/пальца',
        arches: 'свободные',
        diagnosis: 'Беременность ',
        conclusion: conclusion,
        birthPlan: birthPlan
    }
}

function _catchError(e, alerting, cb) {
    console.log(e);
    if (e.error === 'AUTH') {
        if (alerting) {
            alert('Ошибка проверки пользователя.\nПопробуйте выйти/войти в приложение и повторить действие.');
        }
        dispatcher.trigger(dispatcher.RE_LOGIN, {});
    } else if (cb) {
        cb();
    }
}
