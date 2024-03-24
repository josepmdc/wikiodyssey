import React, { useState } from 'react';
import AsyncSelect from 'react-select';

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
}

const InputWithDropdown = (inputWithDropdownProps: InputWithDropdownProps) => {
    const [selectedOption, setSelectedOption] = useState(null);
    const [options, setOptions] = useState([
        { value: 'apple', label: 'Apple' },
        { value: 'banana', label: 'Banana' },
        { value: 'orange', label: 'Orange' }
    ]);

    const handleChange = (selectedOption: any) => {
        if (selectedOption)
            inputWithDropdownProps.inputSelectedCallback(selectedOption.label)
    };

    return (
        <AsyncSelect className='guessInput'
            value={selectedOption}
            onChange={handleChange}
            options={options}
            isClearable
            placeholder="Type your guess..."
        //   styles={customStyles}
        />
    );
};

export default InputWithDropdown;
