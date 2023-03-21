# Walkthrough for Sandbox FE Code

This is a fairly standard React application.

## Directories

`/api`
Api calls that can be made to the backend.

`/apps`
Wraps around a router to serve the application.

`/common`
Common utilities for parsing and helper functions.

`/components`
Shared components used in the application.

`/context`
Application context. See React Context API for more details.

`/layout`
Layout used in the application, including navigation.

`/pages`
Most of the logic for the frontend sits in here. The folder directory tries to map to an available route in `react-router`, much like NextJs.

`/types`
Collection of types used in the application, such as Api references and context interfaces.

## Notable files

`env.ts`
A bunch of environment variables used in the application, like backend api route and the Google client id.
