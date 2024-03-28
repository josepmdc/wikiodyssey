import React, { useEffect, useState } from 'react';
import { ArcherContainer } from 'react-archer';
import './App.css';
import WordElement from './components/WordElement'
import { AnimatePresence } from 'framer-motion';
import InputWithDropdown from './components/InputWithDropdown';
import { getRandomArticles } from './api/wikiodyssey-api';

function App() {
  const [wordChain, setWordChain] = useState<string[]>([]);
  const [attempts, setAttempts] = useState<number>(0)

  const [initialWords, setInitialWords] = useState({
    startingWord: "",
    endingWord: ""
  })

  const addWordToChain = (word: string) => {
    setWordChain(prevChain => [...prevChain, word])
  }

  const addAttempt = () => {
    setAttempts(attempts+1)
  }
  
  useEffect(() => {
    const fetchData = async () => {
      const res = await getRandomArticles(2);
      if(res.articles.length > 1){
        setInitialWords({
          startingWord: res.articles[0],
          endingWord: res.articles[1]
        })
      }
    };
  
    fetchData();
  }, [])
  
  const {startingWord, endingWord} = initialWords

  return (
    <div className="App">
      <header className="App-header">
        <ArcherContainer strokeColor='white' offset={10} key={wordChain.length}>
          <div className='words-container'>
            <AnimatePresence>
              <WordElement key={`initialWord`} content={startingWord} targetId={wordChain.length === 0 ? 'middleWord' : `word0`} elementId='startingWordElement'></WordElement>

              {wordChain.map((word: string, index) => (
                <WordElement
                  key={`${word}-${index}`}
                  content={word}
                  targetId={index === wordChain.length - 1 ? 'middleWord' : `word${index + 1}`}
                  elementId={`word${index}`}></WordElement>
              ))}

              <WordElement key={`middleWord`} content='...' targetId='endingWordElement' elementId='middleWord'></WordElement>
              <WordElement key={`endingWord`} content={endingWord} elementId='endingWordElement'></WordElement>
            </AnimatePresence>
          </div>
        </ArcherContainer>
        
        <InputWithDropdown inputSelectedCallback={addWordToChain} 
                           currentArticleTitle={wordChain.length > 0 ? wordChain[wordChain.length-1] : startingWord}
                           failedAttemptCallback={addAttempt}
                           />

        <div>
          <p>Attempts: {attempts}</p>
        </div>
      </header>
    </div>
  );


}

export default App;
