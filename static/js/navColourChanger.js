/**
 * Updates the background colour of the navigation links in the header based on
 * the current path.
 */

const currentPath = window.location.pathname;

switch (currentPath) {
    case "/":
        updateHomeNav();
        break;

    case "/experience":
        updateExpNav();
        break;

    case "/education":
        updateEduNav();
        break;

    case "/projects":
        updateProjectNav();
        break;

    default:
        break;
}

function updateHomeNav() {
    const navLink = document.getElementById("home-nav");
    changeNavColour(navLink);
}

function updateExpNav() {
    const navLink = document.getElementById("exp-nav");
    changeNavColour(navLink);
}

function updateEduNav() {
    const navLink = document.getElementById("edu-nav");
    changeNavColour(navLink);
}

function updateProjectNav() {
    const navLink = document.getElementById("project-nav");
    changeNavColour(navLink);
}

function changeNavColour(nav) {
    nav.style.background = "gray";
}
