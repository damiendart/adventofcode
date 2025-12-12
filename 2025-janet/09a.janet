# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~{:main (split "\n" (group (* (number :d+) "," (number :d+))))}))

(defn make-rectangles [coordinates]
  (reduce array/concat @[] (map (fn [i] (map (fn [j] (flatten [i j])) coordinates)) coordinates)))

(defn area [rectangle]
  (* (+ (math/abs (- (get rectangle 3) (get rectangle 1))) 1)
     (+ (math/abs (- (get rectangle 0) (get rectangle 2))) 1)))

(defn solve [input]
  (def parsed (peg/match input-parser input))
  (array/pop (sorted (map area (make-rectangles parsed)))))

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
