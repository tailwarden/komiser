import type { Meta, StoryObj } from '@storybook/react';
import ErrorState from './ErrorState';
import mockErrorStateProps from './ErrorState.mocks';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof ErrorState> = {
  title: 'Komiser/Error State',
  component: ErrorState,
  tags: ['autodocs'],
  argTypes: {
    action: {}
  }
};

export default meta;
type Story = StoryObj<typeof ErrorState>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Primary: Story = {
  args: {
    ...mockErrorStateProps.base
  }
};
