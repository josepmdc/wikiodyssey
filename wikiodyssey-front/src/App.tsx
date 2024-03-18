import React, { useState } from 'react';
import { ArcherContainer } from 'react-archer';
import './App.css';
import WordElement from './components/WordElement'

function App() {
  let startingWord = "Banana"
  let endingWord = "Chemistry"

  const [wordChain, setWordChain] = useState<string[]>([]);

  const keyDown = (event: any) => {
    if(event.key === 'Enter'){
      setWordChain(prevChain => [...prevChain, event.currentTarget.value])
      event.target.value = ''
    }
  }

  return (
    <div className="App">
      <header className="App-header">
        <ArcherContainer strokeColor='white' offset={10}>
          <div className='words-container'>

            <WordElement content={startingWord} targetId={wordChain.length === 0 ? 'wordN' : `word0`} elementId='startingWordElement'></WordElement>

            {wordChain.map((word: string, index) => (
              <WordElement 
                key={index} 
                content={word} 
                targetId={index === wordChain.length-1 ? 'wordN' : `word${index+1}`} 
                elementId={`word${index}`}></WordElement>
            ))}

            <WordElement content='...' targetId='endingWordElement' elementId='wordN'></WordElement>
            <WordElement content={endingWord} elementId='endingWordElement'></WordElement>
          </div>
        </ArcherContainer>
        <input className='guessInput' placeholder='Type your guess...' onKeyDown={keyDown}></input>
      </header>
    </div>
  );


}

export default App;
