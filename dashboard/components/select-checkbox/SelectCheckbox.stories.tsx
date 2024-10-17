import type { Meta, StoryObj } from '@storybook/react';
import SelectCheckbox from './SelectCheckbox';
import mockSelectCheckboxProps from './SelectCheckbox.mocks';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof SelectCheckbox> = {
  title: 'Komiser/Select Checkbox',
  component: SelectCheckbox,
  tags: ['autodocs'],
  argTypes: {
    exclude: {},
    setExclude: {}
  }
};

export default meta;
type Story = StoryObj<typeof SelectCheckbox>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Primary: Story = {
  args: {
    ...mockSelectCheckboxProps.base
  }
};
