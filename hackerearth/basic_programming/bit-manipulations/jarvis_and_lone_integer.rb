tests = gets.to_i
tests.times do 
	sum = 0
	n = gets.to_i
	arr = gets.split(" ").map(&:to_i)
	arr.each do |a|
		sum ^= a
	end
	sum = -1 if sum == 0
 
	puts sum
end
