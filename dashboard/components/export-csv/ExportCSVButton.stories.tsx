import type { Meta, StoryObj } from '@storybook/react';
import ExportCSVButton from './ExportCSVButton';

const meta: Meta<typeof ExportCSVButton> = {
  title: 'Komiser/Export CSV',
  component: ExportCSVButton,
  tags: ['autodocs'],
  argTypes: {},
  decorators: [
    Story => (
      <div className="flex items-center justify-center">
        <div className="min-w-24 relative">
          <Story />
        </div>
      </div>
    )
  ]
};

export default meta;
type Story = StoryObj<typeof ExportCSVButton>;

export const Primary: Story = {
  args: {
    disabled: false,
    exportCSV: () => {}
  }
};
