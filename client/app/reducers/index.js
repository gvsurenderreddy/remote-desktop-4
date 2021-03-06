import * as types from '../constants/actions.js'

const initialState = {
	user: {
		name: 'John Doe',
		settings: {
			interface: 'desktop' //'desktop' or 'mobile'
		}
	},

	hosts: {
		avaliable: [],
		current: {
			conn: null, // WebSocket connection with host
			streaming: false, // Host streaming now
			login: null,  // string
		}
	},
	server: {
		addr: NODE_ENV == 'DEVELOP' ? '192.168.0.104:9595' : '46.149.80.62:9595',
	}
}


function user(state = initialState.user, action) {
	switch (action.type) {
		case types.SET_USER:
			return {...state, name: action.name}
		case types.SET_SETTINGS:
			return {...state, settings: {...state.settings, ...action.settings}}
		default:
			return state
	}
}

function hosts(state = initialState.hosts, action) {
	switch (action.type) {
		case types.GET_HOSTS:
			return {...state, avaliable: action.hosts}
		case types.SELECT_HOST:
			if(state.current.conn != null) state.current.conn.close()
			return {...state, current: action.host }
		case types.NEW_FRAME:
			return {...state, current: {...state.current, frame: action.frame} }
		default:
			return state
	}
}

function server(state = initialState.server, action) {
	switch (action.type) {
		default:
			return state
	}
}

const reducers = {
	user,
	hosts,
	server
}

export default reducers



