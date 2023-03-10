import { render, screen } from '@testing-library/react';
import DashboardCloudMap, { DashboardCloudMapProps } from './DashboardCloudMap';

const props: DashboardCloudMapProps = {
  loading: false,
  data: undefined,
  error: false,
  fetch: () => {}
};

describe('Dashboard Cloud Map', () => {
  it('should render without crashing', () => {
    render(<DashboardCloudMap {...props} />);
  });

  it('should render the skeleton component when loading is true', () => {
    render(<DashboardCloudMap {...props} loading={true} />);
    const loading = screen.getByTestId('loading');
    expect(loading).toBeInTheDocument();
  });

  it('should render the error component when error is true', () => {
    render(<DashboardCloudMap {...props} error={true} />);
    const error = screen.getByTestId('error');
    expect(error).toBeInTheDocument();
  });

  it('should render the cloud map card component if error and loading are false', () => {
    render(<DashboardCloudMap {...props} />);
    const cloudMap = screen.getByTestId('cloudMap');
    expect(cloudMap).toBeInTheDocument();
  });
});
