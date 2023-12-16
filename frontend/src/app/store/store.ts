import { combineReducers } from "@ngrx/store";
import { IdentityState, identityReducer } from "./reducer";

export interface IAppState {
  identity: IdentityState
}
