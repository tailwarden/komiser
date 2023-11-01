import { PillProps } from './Pill';

const active: PillProps = {
  status: 'active',
  children: 'active'
};

const pending: PillProps = {
  status: 'pending',
  children: 'pending'
};

const removed: PillProps = {
  status: 'removed',
  children: 'removed'
};

const inactive: PillProps = {
  status: 'inactive',
  children: 'inactive'
};

const info: PillProps = {
  status: 'info',
  children: 'info'
};

const latest: PillProps = {
  status: 'new',
  children: 'latest'
};

const highlight: PillProps = {
  status: 'highlight',
  children: 'highlight'
};

const mockPillProps = {
  active,
  pending,
  removed,
  inactive,
  info,
  latest,
  highlight
};

export default mockPillProps;
