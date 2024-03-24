const apiURL = "http://localhost:1323"

interface getRandomArticlesResponse{
    articles: string[]
}

const getRandomArticles = async (n?: number) => {
    var response = await fetch(`${apiURL}/articles/random?limit=${n ?? 1}`);

    if(!response.ok){
        throw new Error("Random Articles fetching wans't ok");
    }

    const result: getRandomArticlesResponse = await response.json();

    return result;
}


export {getRandomArticles}