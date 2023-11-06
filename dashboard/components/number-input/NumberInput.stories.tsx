import { useState } from 'react';
import type { Meta, StoryObj } from '@storybook/react';
import NumberInput, { InputProps } from './NumberInput';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction

const InputWrapper = ({
  name,
  label,
  value,
  action,
  handleValueChange,
  ...otherProps
}: InputProps) => {
  const [currValue, setCurrValue] = useState<number>(0);
  const handleChange = (newValue: number) => {
    setCurrValue(newValue);
  };
  return (
    <div className="w-96">
      <NumberInput
        name={name}
        label={label}
        value={currValue}
        action={newData => handleChange(Number(newData.title))}
        handleValueChange={handleChange}
        {...otherProps}
      />
    </div>
  );
};

const meta: Meta<typeof NumberInput> = {
  title: 'Komiser/NumberInput',
  component: InputWrapper,
  tags: ['autodocs'],
  argTypes: {
    name: {
      control: 'text',
      description: 'the name for your form (if exist)',
      defaultValue: 'input title'
    },
    label: {
      control: 'text',
      description: 'the label for your input (if exist)',
      defaultValue: ''
    },
    disabled: {
      control: 'boolean',
      description: 'disables the input',
      defaultValue: false
    },
    required: {
      control: 'boolean',
      description: 'Conditionally set the input field as required',
      defaultValue: false
    },
    max: {
      control: 'number',
      description: 'the maximum value'
    },
    min: {
      control: 'number',
      description: 'the minimum value'
    },
    step: {
      control: 'number',
      description: 'change in value',
      defaultValue: false
    },
    maxLength: {
      control: 'number',
      description: 'max length of the input'
    }
  }
};

export default meta;
type Story = StoryObj<typeof NumberInput>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Small: Story = {
  args: {
    name: 'title',
    label: ''
  }
};

export const Large: Story = {
  render: InputWrapper,
  args: {
    name: 'title',
    label: 'Limit'
  }
};
