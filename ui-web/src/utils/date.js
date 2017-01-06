'use strict';

module.exports = {

    parseForInput: (ts) => {
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
        return yyyy + '-' + mm +'-' + dd;
    },

    parseForView: (ts) => {
        if (ts == undefined || ts == null) {
            return '';
        }
        let date = new Date(ts*1000);
        let dd = date.getDate();
        let mm = date.getMonth()+1;
        let yyyy = date.getFullYear();
        if (dd < 10) {
            dd = '0' + dd
        }
        if (mm < 10) {
            mm = '0' + mm
        } 
        return dd + '-' + mm + '-' + yyyy;
      }
};
