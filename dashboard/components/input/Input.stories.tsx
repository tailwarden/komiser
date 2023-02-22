import type { Meta, StoryObj } from '@storybook/react';
import Input from './Input';
import mockInputProps from './Input.mocks';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Input> = {
  title: 'Komiser/Input',
  component: Input,
  tags: ['autodocs'],
  argTypes: {
    regex: {
      control: false
    },
    value: {
      control: false
    },
    type: {
      control: false
    }
  }
};

export default meta;
type Story = StoryObj<typeof Input>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Primary: Story = {
  args: {
    ...mockInputProps.base
  }
};
