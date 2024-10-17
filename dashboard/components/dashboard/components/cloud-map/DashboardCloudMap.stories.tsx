import type { Meta, StoryObj } from '@storybook/react';
import mockDataForDashboard from '../../utils/mockDataForDashboard';
import DashboardCloudMap from './DashboardCloudMap';

const meta: Meta<typeof DashboardCloudMap> = {
  title: 'Komiser Widgets/Cloud Map',
  component: DashboardCloudMap,
  tags: ['autodocs'],
  argTypes: {
    data: {}
  }
};

export default meta;
type Story = StoryObj<typeof DashboardCloudMap>;

export const Loading: Story = {
  args: {
    loading: true,
    error: false
  }
};

export const Error: Story = {
  args: {
    loading: false,
    error: true
  }
};

export const Primary: Story = {
  args: {
    loading: false,
    error: false,
    data: mockDataForDashboard.regions
  }
};
