package com.dudu.daily;

import java.util.*;
import java.util.concurrent.atomic.AtomicInteger;

public class QueenProblem {

    private int[] solution;
    private int size;
    private AtomicInteger solutionCount;

    public QueenProblem(int size) {
        this.size = size;
        solution = new int[size];
        this.solutionCount = new AtomicInteger(0);
    }

    public int[] getSolution() {
        return solution;
    }

    public int getSize() {
        return size;
    }

    public AtomicInteger getSolutionCount() {
        return solutionCount;
    }
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNext()) {
            int size = scanner.nextInt();
            QueenProblem problem = new QueenProblem(size);
            Stack<Point> integerStack = new Stack<>();
            Set<Point> suitableResult = problem.getSuitableResult(0);
            for(Point location : suitableResult){
                integerStack.push(location);
            }
            while (integerStack.size() > 0){
                Point temPoint = integerStack.pop();
                if(temPoint.getRow() == problem.getSize() - 1){
                    problem.getSolutionCount().incrementAndGet();
                    continue;
                }
                problem.modifySolutionStatus(temPoint);
                suitableResult = problem.getSuitableResult(temPoint.getRow() + 1);
                if (suitableResult.size() == 0){
                    continue;
                }
                for(Point location : suitableResult){
                    integerStack.push(location);
                }
            }
            System.out.println(problem.getSolutionCount());
        }
    }

    private void modifySolutionStatus(Point temPoint){
        solution[temPoint.getRow()] = temPoint.getColumn();
        for(int index = temPoint.getRow() + 1; index < size; index++){
            solution[index] = 0;
        }
    }

    static class Point{
        private int row;
        private int column;

        public Point(int row, int column) {
            this.row = row;
            this.column = column;
        }

        public int getRow() {
            return row;
        }

        public void setRow(int row) {
            this.row = row;
        }

        public int getColumn() {
            return column;
        }

        public void setColumn(int column) {
            this.column = column;
        }
    }

    private Set<Point> getSuitableResult(int index) {
        Set<Integer> suitableResult = new HashSet<>();
        for (int i = 0; i < size; i++) {
            suitableResult.add(i);
        }
        for (int i = 0; i < index; i++) {
            suitableResult.remove(solution[i]);
            suitableResult.remove(solution[i] - (index - i));
            suitableResult.remove(solution[i] + (index - i));
        }
        Set<Point> suitResult = new HashSet<>();
        for(int value:suitableResult){
            suitResult.add(new Point(index, value));
        }
        return suitResult;
    }


}
