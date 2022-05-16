import React from 'react'
import {routes, getMountedRoute} from "../../constants/routes";
function Docs() {
  return (
    <>
    <div>
        <a href={getMountedRoute(routes.home)}>
            Home
        </a><br></br>
    </div>
    <h1>Docs</h1>
</>

  )
}

export default Docs;