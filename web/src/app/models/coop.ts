export interface Condition {
	mode: string;
	value: string;
}

export interface Coop {
	opening_condition: Condition;
	closing_condition: Condition;
	latitude: number;
  longitude: number;
  status: string;
}
