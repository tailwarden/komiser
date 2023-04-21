import { EmptyStateProps } from './EmptyState';

const base: EmptyStateProps = {
  title: 'We could not find a cloud account',
  message:
    'It seems you have not connected a cloud account to Komiser. Connect one right now so you can start managing it from your dashboard.',
  actionLabel: 'Guide to connect account',
  mascotPose: 'thinking',
  secondaryActionLabel: 'Report an issue',
  action: () => {},
  secondaryAction: () => {}
};

const mockEmptyStateProps = {
  base
};

export default mockEmptyStateProps;
