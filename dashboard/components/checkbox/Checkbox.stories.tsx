import type { Meta, StoryObj } from '@storybook/react';
import Checkbox from './Checkbox';

const meta: Meta<typeof Checkbox> = {
  title: 'Komiser/Checkbox',
  component: Checkbox,
  tags: ['autodocs']
};

export default meta;
type Story = StoryObj<typeof Checkbox>;

export const Default: Story = {
  args: {
    id: 'checkbox',
    checked: false,
    onChange: () => {}
  }
};

export const Checked: Story = {
  args: {
    id: 'checkbox',
    checked: true,
    onChange: () => {}
  }
};
