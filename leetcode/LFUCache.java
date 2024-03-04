package leetcode;

import java.util.HashMap;
import java.util.LinkedHashSet;
import java.util.Map;

class LFUCache {

    static class Pair {
        int value;
        int freq;

        public Pair(int value) {
            this.value = value;
            this.freq = 1;
        }

    }

    /**
     * key到value+freq的映射
     */
    private Map<Integer, Pair> kp;
    /**
     * 同一频率下的key列表
     */
    private Map<Integer, LinkedHashSet<Integer>> fk;
    /**
     * 最小频率
     */
    private int minFreq;
    /**
     * 容量
     */
    private int capacity;

    public LFUCache(int capacity) {
        kp = new HashMap<>();
        fk = new HashMap<>();
        this.capacity = capacity;
    }

    public int get(int key) {
        if (kp.containsKey(key)) {
            Pair pair = kp.get(key);
            increaseFreq(key, pair);
            return pair.value;
        }
        return -1;
    }

    public void put(int key, int value) {
        if (kp.containsKey(key)) {
            Pair pair = kp.get(key);
            pair.value = value;
            increaseFreq(key, pair);
            return;
        }
        if (kp.size() >= capacity) {
            removeMinFreqKey();
        }
        // 新节点
        Pair pair = new Pair(value);
        fk.computeIfAbsent(1, k -> new LinkedHashSet<>()).add(key);
        kp.put(key, pair);
        this.minFreq = 1;
    }

    private void increaseFreq(int key, Pair pair) {
        // 维护fk 在原先freq删除 添加到新的freq链表尾部
        LinkedHashSet<Integer> oldSet = fk.get(pair.freq);
        // 选LinkedHashSet是因为支持随机删除，hash+双链表
        oldSet.remove(key);
        fk.computeIfAbsent(pair.freq + 1, k -> new LinkedHashSet<>()).add(key);
        if (oldSet.isEmpty()) {
            fk.remove(pair.freq);
            if (minFreq == pair.freq) {
                this.minFreq++;
            }
        }
        pair.freq++;
    }

    private void removeMinFreqKey() {
        LinkedHashSet<Integer> keys = fk.get(minFreq);
        // LinkedHashSet保存了插入顺序，第一个为先插入的元素，即最近未使用的元素
        int key = keys.iterator().next();
        keys.remove(key);
        if (keys.isEmpty()) {
            // 只有一个元素 清理fk
            fk.remove(minFreq);
            // 只有put的时候才会触发剔除，put之后会立马更新 minFreq
        }
        kp.remove(key);
    }
}