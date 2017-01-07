'use strict';

module.exports = {

    parse: (ts, format) => {
        if (ts == undefined || ts == null) {
            return '';
        }
        let date = new Date(ts*1000);
        let dd = date.getDate();
        let mm = date.getMonth()+1; //January is 0!
        let yyyy = date.getFullYear();
        if (dd < 10) {
            dd = '0' + dd
        }
        if (mm < 10) {
            mm = '0' + mm
        }
        switch (format) {
            case 'input':
                return yyyy + '-' + mm +'-' + dd;
            default:
                return dd + '-' + mm + '-' + yyyy;
        }
    }
};
