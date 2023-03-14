import type { Meta, StoryObj } from '@storybook/react';
import ExportCSVButton from './ExportCSVButton';

const meta: Meta<typeof ExportCSVButton> = {
  title: 'Komiser/Export CSV',
  component: ExportCSVButton,
  tags: ['autodocs'],
  argTypes: {}
};

export default meta;
type Story = StoryObj<typeof ExportCSVButton>;

export const Primary: Story = {
  args: {
    loading: false,
    exportCSV: () => {}
  }
};
