import type { Meta, StoryObj } from '@storybook/react';
import Pill from './Pill';
import mockPillProps from './Pill.mocks';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Pill> = {
  title: 'Komiser/Pill',
  component: Pill,
  tags: ['autodocs'],
  argTypes: {}
};

export default meta;
type Story = StoryObj<typeof Pill>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Active: Story = {
  args: {
    ...mockPillProps.active
  }
};

export const Pending: Story = {
  args: {
    ...mockPillProps.pending
  }
};

export const Removed: Story = {
  args: {
    ...mockPillProps.removed
  }
};

export const Inactive: Story = {
  args: {
    ...mockPillProps.inactive
  }
};

export const Info: Story = {
  args: {
    ...mockPillProps.info
  }
};

export const Latest: Story = {
  args: {
    ...mockPillProps.latest
  }
};

export const Highlight: Story = {
  args: {
    ...mockPillProps.highlight
  }
};
