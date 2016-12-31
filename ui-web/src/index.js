'use strict';

import './services/user'
import './services/patient'
import './riotTags'
import dispatcher from './dispatcher';

riot.mount('app');

let changePage = (page, itemId, action) => {
    dispatcher.trigger(dispatcher.PAGE_CHANGE, page, itemId, action);
};
riot.route(changePage);
riot.route.exec(changePage);
riot.route.start();
