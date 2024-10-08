import { render, screen } from '@testing-library/react';
import { RefreshIcon } from '@components/icons';
import Cta from './Cta';

describe('Card', () => {
  it('should render cta component without crashing', () => {
    render(
      <Cta
        title="Introducing Tailwarden"
        description="Tailwarden is the cloud version of Komiser, which offers more features and insights"
        action={<RefreshIcon width={24} height={24} />}
      />
    );
  });

  it('should render the action', () => {
    render(
      <Cta
        title="Introducing Tailwarden"
        description="Tailwarden is the cloud version of Komiser, which offers more features and insights"
        action={<RefreshIcon data-testid="action" width={24} height={24} />}
      />
    );
    const action = screen.getByTestId('action');
    expect(action).toBeInTheDocument();
  });
});
