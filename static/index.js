const DEFAULT_CATEGORY = "Yoshkar-Ola/index";
const GET_PARAM_NAME_CATEGORY = "cat";
const URL_FETCH_NEWS = `fetch-news?${GET_PARAM_NAME_CATEGORY}=`;
const URL_FETCH_CATEGORIES = `fetch-categories`;
const CSS_SELECTED_CLASS = "selected";
const SELECTOR_ID_CATEGORY_POSTFIX = "Category";
const SELECTOR_ID_NEWS_LIST = "newsList";
const SELECTOR_ID_CATEGORIES_LIST = "categoriesList";

function fetchCategories() {
    fetch(URL_FETCH_CATEGORIES)
        .then((response) => {
            return response.text()
        })
        .then((response) => {
            let element = document.getElementById(SELECTOR_ID_CATEGORIES_LIST);
            element.innerHTML = response;

            const currentCategory = getCurrentCategory();
            highlightCategory(currentCategory);

        })
        .catch(() => {
            console.error("Categories fetching error")
        });
}

function fetchNews() {
    const currentCategory = getCurrentCategory();
    fetch(`${URL_FETCH_NEWS}${currentCategory}`)
        .then((response) => {
            return response.text()
        })
        .then((response) => {
            let element = document.getElementById(SELECTOR_ID_NEWS_LIST);
            element.innerHTML = response;
        })

        .catch(() => {
            console.error("News fetching error")
        });
}

function getCurrentCategory() {
    const currentUrlString = window.location.href;
    const url = new URL(currentUrlString);
    let category = url.searchParams.get(GET_PARAM_NAME_CATEGORY);
    return category ? category : DEFAULT_CATEGORY;
}

function highlightCategory(cat) {
    const catElement = document.getElementById(cat + SELECTOR_ID_CATEGORY_POSTFIX);
    if (catElement) {
        catElement.classList.add(CSS_SELECTED_CLASS);
    } else {
        console.error(`Unknown category ${cat}`);
    }
}

fetchCategories();
fetchNews();


