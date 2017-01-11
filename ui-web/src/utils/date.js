'use strict';

import moment from 'moment';

module.exports = {

    parse: (from) => {
        return moment(from, 'DD-MM-YYYY HH:mm');
    },
    format: (ts) => {
        if (ts == undefined || ts == null) {
            return '';
        }
        return moment.unix(ts).format('DD-MM-YYYY HH:mm');
    }
};
