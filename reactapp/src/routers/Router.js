import React from "react";

import { BrowserRouter as Router, Switch } from "react-router-dom";
import Home from "../components/home";
import Docs from "../components/docs";

import CommonRoute from "./CommonRoute";
import {routes, getMountedRoute} from "../constants/routes";

function AppRouter() {
  return (
    <Router>
      <Switch>
        <CommonRoute component={Docs} path={getMountedRoute(routes.docs.root)} />
        <CommonRoute component={Home} path={getMountedRoute(routes.home)} exact={true} />
      </Switch>
    </Router>
  );
}

export default AppRouter;