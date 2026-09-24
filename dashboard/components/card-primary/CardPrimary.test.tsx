import { render, screen } from '@testing-library/react';
import CardPrimary from './CardPrimary';

describe('Card', () => {
  it('should render cta component without crashing', () => {
    render(
      <CardPrimary
        title="Introducing Tailwarden"
        description="Tailwarden is the cloud version of Komiser, which offers more features and insights"
        type="shadow"
      />
    );
  });
});
