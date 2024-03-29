const apiURL = "http://localhost:1323"

interface GetRandomArticlesResponse{
    articles: string[]
}

interface GetTitleSuggestionsResponse{
    titles: {
        description: string
        title: string
        id: number
        key: string
    }[]
}

const getRandomArticles = async (n?: number) => {
    var response = await fetch(`${apiURL}/articles/random?limit=${n ?? 1}`);

    if(!response.ok){
        throw new Error("Random Articles fetching wasn't ok");
    }

    const result: GetRandomArticlesResponse = await response.json();

    return result;
}

interface Option {
    value: string
    label: string
}

function convertTitlesToOptions(response: GetTitleSuggestionsResponse): any{
    var titles: Option[] = []
    response.titles.forEach((element) => {
        titles.push({
            value: element.key,
            label: element.title
        })
    });

    return titles
}

const getTitleSuggestions = async (input: string) : Promise<Option[]> => {
    var response = await fetch(`${apiURL}/articles/titles?input=${input}`);

    if(!response.ok){
        throw new Error("Title suggestions fetching wasn't ok");
    }

    const result: GetTitleSuggestionsResponse = await response.json();
    
    return convertTitlesToOptions(result);
}

const getIsTitleInArticle = async (sourceTitle: string, targetTitle: string) : Promise<string> => {
    var response = await fetch(`${apiURL}/articles/IsTitleInArticle?sourceTitle=${sourceTitle}&targetTitle=${targetTitle}`);

    if(!response.ok){
        throw new Error("Is title in article fetching wasn't ok");
    }

    const result: string = await response.json();
    
    return result;
}


export {getRandomArticles, getTitleSuggestions, getIsTitleInArticle}