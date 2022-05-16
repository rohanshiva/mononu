import React from 'react'
import {routes, getMountedRoute} from "../../constants/routes";
function Home() {
    return (
        <>
            <div>
                <a href={getMountedRoute(routes.docs.root)}>
                    Docs
                </a>
            </div>
            <h1>Mononu home page</h1>

        </>

    )
}

export default Home;