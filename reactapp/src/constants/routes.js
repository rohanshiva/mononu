const MOUNTED_PATH = "reactapp"

function getMountedRoute(route) {
    return `/${MOUNTED_PATH}${route}`;
}

const routes = {
    "home": "/",
    "docs": {
        "root": "/docs",
    },
}

export {routes, getMountedRoute, MOUNTED_PATH};
