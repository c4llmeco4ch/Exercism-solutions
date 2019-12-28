def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("Lengths of A and B must be equal")
    return sum([1 for a_val, b_val in zip(strand_a,strand_b) if a_val != b_val])
