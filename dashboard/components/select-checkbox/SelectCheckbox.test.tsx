import { render, fireEvent } from '@testing-library/react';
import SelectCheckbox from './SelectCheckbox';

jest.mock('./hooks/useSelectCheckbox', () => ({
  __esModule: true,
  default: jest.fn(() => ({
    listOfExcludableItems: ['Item 1', 'Item 2', 'Item 3'],
    error: null
  }))
}));

describe('SelectCheckbox', () => {
  it('renders the label correctly', () => {
    const { getByText } = render(
      <SelectCheckbox
        label="Test Label"
        query="provider"
        exclude={[]}
        setExclude={() => {}}
      />
    );

    expect(getByText('Test Label')).toBeInTheDocument();
  });

  it('opens the dropdown when clicked', () => {
    const { getByRole, getByText } = render(
      <SelectCheckbox
        label="Test Label"
        query="provider"
        exclude={[]}
        setExclude={() => {}}
      />
    );

    fireEvent.click(getByRole('button'));

    expect(getByText('Item 1')).toBeInTheDocument();
    expect(getByText('Item 2')).toBeInTheDocument();
    expect(getByText('Item 3')).toBeInTheDocument();
  });

  it('closes the dropdown when clicked outside', () => {
    const { getByRole, queryByText, getByTestId } = render(
      <SelectCheckbox
        label="Test Label"
        query="provider"
        exclude={[]}
        setExclude={() => {}}
      />
    );

    fireEvent.click(getByRole('button'));
    expect(queryByText('Item 1')).toBeInTheDocument();

    fireEvent.click(getByTestId('overlay'));
    expect(queryByText('Item 1')).toBeNull();
  });

  it('selects and excludes items', () => {
    const setExcludeMock = jest.fn();
    const { getByRole, getByText } = render(
      <SelectCheckbox
        label="Test Label"
        query="provider"
        exclude={[]}
        setExclude={setExcludeMock}
      />
    );

    fireEvent.click(getByRole('button'));
    fireEvent.click(getByRole('checkbox', { name: 'Item 1' }));
    fireEvent.click(getByRole('checkbox', { name: 'Item 3' }));
    fireEvent.click(getByText('Apply'));

    expect(setExcludeMock).toHaveBeenCalledWith(['Item 1', 'Item 3']);
  });

  it('selects and deselects all items with "Exclude All" checkbox', () => {
    const setExcludeMock = jest.fn();
    const { getByRole, getByText } = render(
      <SelectCheckbox
        label="Test Label"
        query="provider"
        exclude={[]}
        setExclude={setExcludeMock}
      />
    );

    fireEvent.click(getByRole('button'));
    fireEvent.click(getByRole('checkbox', { name: 'Exclude All' }));
    fireEvent.click(getByText('Apply'));

    expect(setExcludeMock).toHaveBeenCalledWith(['Item 1', 'Item 2', 'Item 3']);

    fireEvent.click(getByRole('checkbox', { name: 'Exclude All' }));
    fireEvent.click(getByText('Apply'));

    expect(setExcludeMock).toHaveBeenCalledWith([]);
  });

  it('filters the list of items based on search input', () => {
    const { getByRole, getByPlaceholderText, queryByText } = render(
      <SelectCheckbox
        label="Test Label"
        query="provider"
        exclude={[]}
        setExclude={() => {}}
      />
    );

    fireEvent.click(getByRole('button'));
    const searchInput = getByPlaceholderText('Search');

    fireEvent.change(searchInput, { target: { value: 'Item 1' } });
    expect(queryByText('Item 2')).toBeNull();
    expect(queryByText('Item 3')).toBeNull();
    expect(queryByText('Item 1')).toBeInTheDocument();

    fireEvent.change(searchInput, { target: { value: 'Item' } });
    expect(queryByText('Item 1')).toBeInTheDocument();
    expect(queryByText('Item 2')).toBeInTheDocument();
    expect(queryByText('Item 3')).toBeInTheDocument();

    fireEvent.change(searchInput, { target: { value: 'Non-matching' } });
    expect(queryByText('Item 1')).toBeNull();
    expect(queryByText('Item 2')).toBeNull();
    expect(queryByText('Item 3')).toBeNull();
  });
});
