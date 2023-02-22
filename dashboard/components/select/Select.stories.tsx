import type { Meta, StoryObj } from '@storybook/react';
import Select from './Select';
import mockSelectProps from './Select.mocks';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Select> = {
  title: 'Komiser/Select',
  component: Select,
  tags: ['autodocs'],
  argTypes: {}
};

export default meta;
type Story = StoryObj<typeof Select>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Primary: Story = {
  args: {
    ...mockSelectProps.base
  }
};
