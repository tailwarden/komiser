import type { Meta, StoryObj } from '@storybook/react';
import Template from './Template';
import mockTemplateProps from './Template.mocks';

// More on how to set up stories at: https://storybook.js.org/docs/7.0/react/writing-stories/introduction
const meta: Meta<typeof Template> = {
  title: 'Komiser/Template',
  component: Template,
  tags: ['autodocs'],
  argTypes: {}
};

export default meta;
type Story = StoryObj<typeof Template>;

// More on writing stories with args: https://storybook.js.org/docs/7.0/react/writing-stories/args
export const Primary: Story = {
  args: {
    ...mockTemplateProps.base
  }
};
