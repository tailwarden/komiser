import type { Meta, StoryObj } from '@storybook/react';
import Avatar from './Avatar';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Avatar> = {
  title: 'Komiser/Avatar',
  component: Avatar,
  tags: ['autodocs']
};

export default meta;
type Story = StoryObj<typeof Avatar>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const CloudProviders: Story = {
  args: {
    avatarName: 'aws'
  }
};

export const Integrations: Story = {
  args: {
    avatarName: 'slack'
  }
};
