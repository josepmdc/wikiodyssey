import React, { useState } from 'react';
import AsyncSelect from 'react-select/async';
import { getIsTitleInArticle, getTitleSuggestions } from '../api/wikiodyssey-api';

const customStyles = {
    control: (provided: any, state: { isFocused: any; }) => ({
        ...provided,
        width: '40%',
        minHeight: '80px', // Adjust the height as needed
        fontSize: 'xx-large',
        borderRadius: '20px',
        color: 'rgb(65, 62, 62)',
        textAlign: 'center',
        border: state.isFocused ? '1px solid #2684ff' : '1px solid #ccc',
        boxShadow: state.isFocused ? '0 0 0 1px #2684ff' : null,
        '&:hover': {
            borderColor: '#2684ff',
        },
    }),
    menu: (provided: any, state: { isFocused: any; }) => ({
        ...provided,
        margin: '0px'
    }),
};

interface InputWithDropdownProps {
    inputSelectedCallback: Function
    currentArticleTitle: string
    failedAttemptCallback: Function
}

interface GetTitleSuggestionsResponse{
    titles: {
        description: string
        title: string
        id: number
    }[]
}



const InputWithDropdown = (inputWithDropdownProps: InputWithDropdownProps) => {
    const [selectedOption, setSelectedOption] = useState(null);

    const handleChange = async (selectedOption: any) => {
        if (selectedOption){
            setSelectedOption(null)
            console.log("Current article title is: ", inputWithDropdownProps.currentArticleTitle)

            try{
                const response = await getIsTitleInArticle(inputWithDropdownProps.currentArticleTitle, selectedOption.value)
                console.log("Response: ", response)
                inputWithDropdownProps.inputSelectedCallback(selectedOption.label)
            }catch(exception){
                console.log(exception)
                inputWithDropdownProps.failedAttemptCallback()
            }

        }
    };

    const loadOptions = (inputValue: string, callback: any) => {
        if(inputValue.length > 0){
            var response = getTitleSuggestions(inputValue)
            return response
        }
    }

    return (
        <AsyncSelect className='guessInput'
            value={selectedOption}
            onChange={handleChange}
            loadOptions={loadOptions}
            // onInputChange={handleInputChange}
            // options={options}
            isClearable
            placeholder="Type your guess..."
        //   styles={customStyles}
        />
    );
};

export default InputWithDropdown;
