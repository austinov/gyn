'use strict';

module.exports = {
    // Helper method to check response status in fetch.
    handleStatus: (response, isBlob) => {
        let s = response.status, e;
        if (s === 200 || s === 0) {
            if (isBlob) {
                let filename = "";
                let disposition = response.headers.get('Content-Disposition');
                if (disposition && disposition.indexOf('attachment') !== -1) {
                    let filenameRegex = /filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/;
                    let matches = filenameRegex.exec(disposition);
                    if (matches != null && matches[1]) filename = matches[1].replace(/['"]/g, '');
                }
                return Promise.resolve(response.blob())
                    .then(blob => {
                        return Promise.resolve({
                            filename: filename,
                            blob: blob
                        });
                    });
            } else {
                return Promise.resolve(response.json());
            }
        } else if (s === 401 || s === 403) {
            e = 'AUTH';
        } else {
            e = 'ERROR';
        }
        let generalHandler = err => {
            let msg = 'HTTP error: ' +
                response.statusText + '.';
            return Promise.reject({
                error: e,
                msg: msg,
                cause: new Error(response.statusText)
            });
        }
        if (response.headers.get('Content-Type').startsWith('application/json')) {
            return Promise.resolve(response.json())
                .then(
                    err => {
                        return Promise.reject({
                            error: e,
                            msg: msgForErr(err),
                            cause: new Error(response.statusText)
                        });
                },
                generalHandler);
        }
        return Promise.resolve(response.text())
            .then(
                err => {
                    let msg = 'HTTP error: ' +
                        response.statusText + ', ' +
                        err + '.';
                    return Promise.reject({
                        error: e,
                        msg: msg,
                        cause: new Error(response.statusText)
                    });
            },
            generalHandler);
    }
}

// Sample function for error message customization.
// Error code could be mapped to localized message using some localization lib
// (like i18next). Here, for simplicity, we just provide more readable English text.
function msgForErr(err) {
    if (Array.isArray(err)) {
        let msg = '';
        for (let i = 0, l = err.length; i < l; i++) {
            msg += msgForErr(err[i]);
            msg += ' ';
        }
        return msg;
    }
    if (!err.code || !err.message) {
        return err;
    }
    switch (err.code) {
		case 'login-required':
		case 'login-format':
		case 'password-required':
		case 'password-format':
		case 'no_data_err':
		case 'srv_err':
            return err.message + '.';
        case 'invalid_req_param':
            return 'Invalid request parameter.';
        case 'auth_err':
            return 'Incorrect username or password.';
    }
    return 'Server reported unrecognized error: ' + err.message;
}
