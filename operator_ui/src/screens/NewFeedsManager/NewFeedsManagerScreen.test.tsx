import '@testing-library/jest-dom'

import * as React from 'react'
import {
  render,
  screen,
  waitFor,
  waitForElementToBeRemoved,
} from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { MockedProvider, MockedResponse } from '@apollo/client/testing'

import { FETCH_FEEDS_MANAGERS } from 'src/hooks/useFetchFeedsManager'
import {
  CREATE_FEEDS_MANAGER,
  NewFeedsManagerScreen,
} from './NewFeedsManagerScreen'
import { RedirectTestRoute } from 'support/test-helpers/RedirectTestRoute'

import { MemoryRouter, Route } from 'react-router-dom'
import { Provider } from 'react-redux'
import createStore from 'createStore'
import { buildFeedsManager } from 'support/test-helpers/factories/feedsManager'

test('renders the page', async () => {
  const mocks: MockedResponse[] = [
    {
      request: {
        query: FETCH_FEEDS_MANAGERS,
      },
      result: {
        data: {
          feedsManagers: {
            results: [],
          },
        },
      },
    },
  ]

  render(
    <MockedProvider mocks={mocks} addTypename={false}>
      <Provider store={createStore()}>
        <MemoryRouter>
          <NewFeedsManagerScreen />
        </MemoryRouter>
      </Provider>
    </MockedProvider>,
  )

  await waitForElementToBeRemoved(() => screen.queryByRole('progressbar'))

  expect(screen.queryByText('Register Feeds Manager')).toBeInTheDocument()
  expect(screen.queryByTestId('feeds-manager-form')).toBeInTheDocument()
})

test('redirects when a manager exists', async () => {
  const mocks: MockedResponse[] = [
    {
      request: {
        query: FETCH_FEEDS_MANAGERS,
      },
      result: {
        data: {
          feedsManagers: {
            results: [buildFeedsManager()],
          },
        },
      },
    },
  ]

  render(
    <MockedProvider mocks={mocks} addTypename={false}>
      <Provider store={createStore()}>
        <MemoryRouter>
          <Route exact path="/">
            <NewFeedsManagerScreen />
          </Route>

          <RedirectTestRoute
            path="/feeds_manager"
            message="Feeds Manager Page"
          />
        </MemoryRouter>
      </Provider>
    </MockedProvider>,
  )

  await waitForElementToBeRemoved(() => screen.queryByRole('progressbar'))

  expect(screen.queryByText('Feeds Manager Page')).toBeInTheDocument()
})

test('submits the form', async () => {
  const { getByRole, getByTestId } = screen
  const mocks: MockedResponse[] = [
    {
      request: {
        query: FETCH_FEEDS_MANAGERS,
      },
      result: {
        data: {
          feedsManagers: {
            results: [],
          },
        },
      },
    },
    {
      request: {
        query: CREATE_FEEDS_MANAGER,
        variables: {
          input: {
            name: 'Chainlink Feeds Manager',
            uri: 'localhost:8080',
            publicKey: '1111',
            jobTypes: ['FLUX_MONITOR'],
            isBootstrapPeer: false,
            bootstrapPeerMultiaddr: undefined,
          },
        },
      },
      result: {
        data: {
          createFeedsManager: {
            feedsManager: buildFeedsManager(),
          },
        },
      },
    },
    {
      request: {
        query: FETCH_FEEDS_MANAGERS,
      },
      result: {
        data: {
          feedsManagers: {
            results: [buildFeedsManager()],
          },
        },
      },
    },
  ]

  render(
    <MockedProvider mocks={mocks} addTypename={false}>
      <Provider store={createStore()}>
        <MemoryRouter>
          <Route exact path="/">
            <NewFeedsManagerScreen />
          </Route>

          <RedirectTestRoute
            path="/feeds_manager"
            message="Feeds Manager Page"
          />
        </MemoryRouter>
      </Provider>
    </MockedProvider>,
  )

  await waitForElementToBeRemoved(() => screen.queryByRole('progressbar'))

  // Note: The name input has a default value so we don't have to set it
  userEvent.type(getByRole('textbox', { name: 'URI *' }), 'localhost:8080')
  userEvent.type(getByRole('textbox', { name: 'Public Key *' }), '1111')
  userEvent.click(getByRole('checkbox', { name: 'Flux Monitor' }))

  userEvent.click(getByTestId('create-submit'))

  await waitFor(() =>
    expect(screen.queryByText('Feeds Manager Page')).toBeInTheDocument(),
  )
})
