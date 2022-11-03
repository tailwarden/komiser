import { EmptyStateProps } from './EmptyState';

const base: EmptyStateProps = {
  title: 'Add a cloud account',
  message: 'Connect your cloud accounts and uncover hidden costs with Oraculi.',
  action: () => {},
  actionLabel: 'Add new account',
  mascotPose: 'greetings'
};

const mockEmptyStateProps = {
  base
};

export default mockEmptyStateProps;
