
 

import API from './API';
import { mockData } from '../___Mocks___/MockData/MockData';
import axios, { AxiosInstance } from 'axios';
import { Container } from '../interfaces/container.interface';
import moxios from 'moxios';

let testContainer: any;

beforeEach(() => {
  testContainer = mockData[0];
});

describe('API', () => {
  const axiosInstance = new API();
  beforeEach(() => {
    moxios.install(axiosInstance._baseUrl);
  });

  afterEach(() => {
    moxios.uninstall(axiosInstance._baseUrl);
  });

  it('should get all containers', async () => {
    moxios.wait(() => {
      const request = moxios.requests.mostRecent();
      request.respondWith({ status: 200, response: mockData });
    });
    const response = await axiosInstance.getAllContainers();
    expect(response.data).toEqual(mockData);
  });

  it('should get a specific container', async () => {
    moxios.wait(() => {
      const request = moxios.requests.mostRecent();
      request.respondWith({ status: 200, response: mockData[0] });
    });
    const response = await axiosInstance.getContainer(0);
    expect(response.data).toEqual(mockData[0]);
  });

  it('should update a container', async () => {
    moxios.wait(() => {
      const request = moxios.requests.mostRecent();
      request.respondWith({ status: 200, response: mockData[0] });
    });
    const response = await axiosInstance.updateContainer(mockData[0]);
    expect(response.data).toEqual(mockData[0]);
  });
});