import React, { useState } from 'react';
import { ArcherContainer } from 'react-archer';
import './App.css';
import WordElement from './components/WordElement'
import { AnimatePresence } from 'framer-motion';

function App() {
  let startingWord = "Banana"
  let endingWord = "Chemistry"

  const [wordChain, setWordChain] = useState<string[]>([]);

  const keyDown = (event: any) => {
    if (event.key === 'Enter' && event.target.value !== '') {
      setWordChain(prevChain => [...prevChain, event.currentTarget.value])
      event.target.value = ''
    }
  }
  
  return (
    <div className="App">
      <header className="App-header">
        <ArcherContainer strokeColor='white' offset={10} key={wordChain.length}>
          <div className='words-container'>
            <AnimatePresence>
              <WordElement key={`initialWord-${new Date().getTime()}`} content={startingWord} targetId={wordChain.length === 0 ? 'middleWord' : `word0`} elementId='startingWordElement'></WordElement>

              {wordChain.map((word: string, index) => (
                <WordElement
                  key={`${word}-${index}-${new Date().getTime()}`}
                  content={word}
                  targetId={index === wordChain.length - 1 ? 'middleWord' : `word${index + 1}`}
                  elementId={`word${index}`}></WordElement>
              ))}

              <WordElement key={`middleWord-${new Date().getTime()}`} content='...' targetId='endingWordElement' elementId='middleWord'></WordElement>
              <WordElement key={`endingWord-${new Date().getTime()}`} content={endingWord} elementId='endingWordElement'></WordElement>
            </AnimatePresence>
          </div>
        </ArcherContainer>
        <input className='guessInput' placeholder='Type your guess...' onKeyDown={keyDown}></input>
      </header>
    </div>
  );


}

export default App;
