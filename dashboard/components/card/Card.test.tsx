import { render, screen } from '@testing-library/react';
import RefreshIcon from '../icons/RefreshIcon';
import Card from './Card';

describe('Card', () => {
  it('should render card component without crashing', () => {
    render(
      <Card
        label="Test card"
        value={500}
        icon={<RefreshIcon width={24} height={24} />}
      />
    );
  });

  it('should render the icon', () => {
    render(
      <Card
        label="Test card"
        value={500}
        icon={<RefreshIcon width={24} height={24} />}
      />
    );
    const icon = screen.getByTestId('icon');
    expect(icon).toBeInTheDocument();
  });

  it('should display the value formatted', () => {
    render(
      <Card
        label="Test card"
        value={5000}
        icon={<RefreshIcon width={24} height={24} />}
      />
    );
    const formattedNumber = screen.getByTestId('formattedNumber');
    expect(formattedNumber).toHaveTextContent('5K');
  });

  it('should display the value formatted as currency', () => {
    render(
      <Card
        label="Test card"
        value={6000}
        icon={<RefreshIcon width={24} height={24} />}
        formatter="currency"
      />
    );
    const formattedNumber = screen.getByTestId('formattedNumber');
    expect(formattedNumber).toHaveTextContent('$6K');
  });

  it('should render tooltip if there is a tooltip set', () => {
    render(
      <Card
        label="Test card"
        value={500}
        tooltip="This will be the tooltip text"
        icon={<RefreshIcon width={24} height={24} />}
      />
    );
    const tooltip = screen.getByRole('tooltip');
    expect(tooltip).toHaveTextContent('This will be the tooltip text');
  });
});
